#!/vault-home/goelayu/build/bin/Rscript
# Arguments: <path to input file> <path to output file>
library(ggplot2)
library(wesanderson)

args <- commandArgs(trailingOnly=TRUE)
output <- args[2]

data <- read.csv(args[1], header=TRUE, sep=" ")

# data$network <- factor(data$network, levels=c('2000', '2010', '2020'))
# pdf(output, height=2.5, width=5.0)
png(output, height=2.5, width=5.0, units="in", res=300)
ggplot(data, aes(x=x,color=factor(y) )) +
	stat_ecdf(pad = FALSE,size=1) +
	# geom_ribbon(aes(ymin=plt, ymax=1)) + 
	#geom_line()+
	xlab(" # resources fetched ") +
	# xlab("(Main image bytes)/(Total image bytes)") +
	ylab("CDF across news sites") +
	# ylab(expression(atop("CDF across pages"))) + 
	guides(fill=guide_legend(title="Config")) +
	coord_cartesian(ylim=c(0, 1),xlim=c(0,180)) +
	scale_x_continuous(expand=c(0, 0)) +
	scale_y_continuous(expand=c(0, 0)) + 
	 # scale_color_brewer(breaks=breaks, limits=breaks) + 
	 # scale_x_continuous(sec.axis = sec_axis(~.*1000000, name = "Unique Invocation Count")) + 
	# scale_linetype_manual(values=c("dotted","solid","longdash")) + 
	theme_bw() +
	# scale_linetype_manual(values = c("solid","dash","longdash")) + 
	# guides(fill = guide_legend(nrow = 1))+
	# guides(colour=guide_legend(nrow=2, byrow=TRUE), linetype=guide_legend(nrow=2, byrow=TRUE), direction="vertical") +
	theme(
		legend.title=element_blank(),
		legend.background=element_rect(color="darkgray", fill="white", linetype="solid", size=0.3),
		legend.key=element_blank(),
		legend.key.height=unit(12, "points"),
		legend.key.width=unit(30, "points"),
		legend.position=c(0.72, 0.33),
		# legend.position="top",
		legend.margin=margin(c(1,3,3,3)),
		axis.title=element_text(size=14),
		axis.text=element_text(size=14),
		legend.text=element_text(size=14),
		axis.title.y=element_text(margin=margin(0, 0, 0, 0)),
		# axis.text.x=element_text(margin=margin(10, 0, 0, 0),angle=45),
		axis.text.x=element_text(margin=margin(10, 0, 0, 0)),
		plot.margin=unit(c(15,25,5,15),"points"))
		


.junk <- dev.off()

# lines(ecdf(data2$plt))
#for the js to computation ratio, the input file used was js_fraction_cold.csv