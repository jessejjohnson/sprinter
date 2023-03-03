#!/usr/bin/env bash

# Extracts list of URLs and generates an input file
# for the chrome-distributed crawler

# Usage: ./run.sh <sites_file> <pages_dst> <output> <proxypath> <args> <nCrawlers>
# sites_file: file containing list of URLs
# pages_dst: directory to store downloaded pages
# output: output directory
# proxypath: path to recorded pages
# args: arguments to pass to crawler
# nCrawlers: number of crawlers to spawn

# Example: ./run.sh sites.txt ../../pageshere ../data/output `../data/record/output` "-t 10" 30

# ensure ulimit is large enough
# sudo ulimit -n 100000

trap ctrl_c SIGINT

function ctrl_c() {
    echo "** Trapped CTRL-C"
    echo "** Stopping all processes"
    ps aux | grep sys-usage-track | awk '{print $2}' | xargs kill -9
}

echo "number of arguments: $#"
echo "arguments: $@"
if [ $# -ne 6 ]; then
    echo "Usage: ./run.sh <sites_file> <pages_dst> <output> <wprpath> <args>"
    echo "sites_file: file containing list of URLs"
    echo "pages_dst: directory containing pages per site"
    echo "rundir: run directory -- stored output and copies wprdata to this"
    echo "wprpath: path to recorded pages"
    echo "args: arguments to pass to crawler"
    exit 1
fi

NPAGES=10
CHROMESCRIPT=../node/chrome-distributed.js
WPRDATA=/w/goelayu/bcrawling/wprdata
LOGDIR=../../logs/eval
tmpdir=/run/user/99542426/goelayu/tmp
mkdir -p $tmpdir

sites_file=$1
pages_dst=$2
rundir=$3
wprpath=$4
args=$5

urlfile=$tmpdir/urls.txt
infile=$tmpdir/infile.txt

cat $sites_file > $infile

create_url_file(){
    rm -f $urlfile
    for i in $(seq 1 $NPAGES); do
        while read site; do
            p=`sed -n ${i}p $pages_dst/${site}.txt`;
            echo $site $p >> $urlfile
        done < $infile
    done
}

create_url_dist(){
    pagesfile=$1
    ncrawlers=$2
    NPAGES=`cat $pagesfile | wc -l`
    
    disturldir=$tmpdir/dist
    mkdir -p $disturldir
    
    pagepercrawl=$((NPAGES / ncrawlers))
    t=$pagepercrawl
    h=0
    for i in $(seq 1 $ncrawlers); do
        h=$((h + pagepercrawl))
        urlfile=$disturldir/urls-$i.txt
        rm -f $urlfile
        cat $pagesfile | head -n $h | tail -n $t | while read site; do
            cat $pages_dst/${site}.txt >> $urlfile
        done
    done
}

if [[ "$COPY" == "1" ]]; then
    echo "Copying wprdata"
    while read line; do
        site=`echo $line | cut -d' ' -f1`
        page=`echo $line | cut -d' ' -f2 | sanitize`
        cp $wprpath/$site/$page/data.wprgo $WPRDATA/${page}.wprgo 2>/dev/null &
    done < $urlfile
    wait
    echo "Done copying wprdata"
else
    echo "Not copying wprdata"
fi

# sys usage tracking
sysscript=../../bash/profiling/sys-usage-track.sh
mkdir -p $rundir/sys
rm -rf $rundir/sys/*
$sysscript $rundir/sys &
sysupid=$!;

mkdir -p $rundir/output
# rm -rf $rundir/output/*

# copy the static analysis tool to tmpfs
cp -r /vault-swift/goelayu/balanced-crawler/node/program_analysis/ /run/user/99542426/goelayu/panode/

echo "Starting the az server"
# start the az server
GOROOT="/w/goelayu/uluyol-sigcomm/go";
AZDIR=/vault-swift/goelayu/balanced-crawler/system/go/wpr
AZPORT=`shuf -i 8000-16000 -n 1`
(cd $AZDIR; GOGC=off GOROOT=${GOROOT} go run src/analyzer/main.go src/analyzer/rewriter.go src/analyzer/genjs.go --port $AZPORT &> $rundir/output/az.log ) &
# azpid=$!

create_crawl_instances_baseline(){
    ncrawlers=$1
    PERSCRIPTCRAWLERS=5
    nscripts=$((ncrawlers / PERSCRIPTCRAWLERS))
    echo "Creating $nscripts scripts"
    first=0
    for i in $(seq 1 $nscripts); do
        totalurls=`cat $urlfile | wc -l`;
        scripturls=`echo $totalurls / $nscripts | bc`;
        first=$((first + scripturls));
        echo "Running cmd: node $CHROMESCRIPT -u <(cat $urlfile | awk '{print \$2}' | head -n $first | tail -n $scripturls)\
        -o $rundir/output --proxy $WPRDATA $args --azport $AZPORT -c $PERSCRIPTCRAWLERS"
        { time node $CHROMESCRIPT -u <(cat $urlfile | awk '{print $2}' | head -n $first | tail -n $scripturls) -o $rundir/output \
        --proxy $WPRDATA $args --azport $AZPORT -c $PERSCRIPTCRAWLERS ; } &> $LOGDIR/$LOGFILE-$i.log &
        pids[${i}]=$!
    done
    echo "Waiting for all scripts to finish"
    for pid in ${pids[*]}; do
        wait $pid
    done
    echo "All scripts finished"
}

create_crawl_instances_opt(){
    ncrawlers=$1
    PERSCRIPTCRAWLERS=1
    for i in $(seq 1 $ncrawlers); do
        [ ! -f $tmpdir/dist/urls-$i.txt ] && echo "File $tmpdir/dist/urls-$i.txt does not exist" && exit 1
        
        urlfile=$tmpdir/dist/urls-$i.txt
        echo "Running cmd: node $CHROMESCRIPT -u $urlfile -o $rundir/output --proxy $WPRDATA $args --azport $AZPORT -c $PERSCRIPTCRAWLERS &> $LOGDIR/$LOGFILE-$i.log "
        { time node $CHROMESCRIPT -u $urlfile -o $rundir/output --proxy $WPRDATA $args --azport $AZPORT -c $PERSCRIPTCRAWLERS ; } &> $LOGDIR/$LOGFILE-$i.log &
        pids[${i}]=$!
    done
    echo "Waiting for all scripts to finish"
    for pid in ${pids[*]}; do
        wait $pid
    done
}

if [[ "$RUN" == "baseline" ]]; then
    echo "Creating the default url file"
    create_url_file
    echo "Starting the baseline crawling script"
    create_crawl_instances_baseline $6
    elif [[ "$RUN" == "opt" ]]; then
    echo "Creating the optimized url file"
    create_url_dist $infile $6
    echo "Starting the optimized crawling script"
    create_crawl_instances_opt $6
else
    echo "Invalid run type"
fi


# kill the resource usage scripts
echo "Sending ctrl-c to the monitoring tools" $sysupid
kill -SIGUSR1 $sysupid;
# kill all usage scripts
ps aux | grep usage | grep -v grep | awk '{print $2}' | xargs kill -9

# kill the az server
ps aux | grep $AZPORT | grep -v grep | awk '{print $2}' | xargs kill -SIGINT