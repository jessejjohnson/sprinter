// Script invoked by the go rewriter to instrument
// all the javaScript on a given web page

const fs = require("fs");
const program = require("commander");
const stateTracker = require("./static/track-file-state.js");
const DYNPATH =
  "/vault-swift/goelayu/balanced-crawler/node/program_analysis/dynamic/tracer.js";
const htmlparser = require("node-html-parser");

program
  .version("0.0.1")
  .option("-i, --input [input]", "The input file")
  .option("-t, --type [type]", "The type of file to instrument")
  .option("-n, --name [name]", "The name of the instrumented file")
  .option(
    "-f, --identifier [identifier]",
    "The identifier of the instrumented file"
  )
  .option("--analyzing [analyzing]", "Whether to analyze the file or not")
  .parse(process.argv);

if (!program.input) {
  console.log("Please specify an input file");
  process.exit(1);
}

function IsJsonString(str) {
  try {
    JSON.parse(str);
  } catch (e) {
    return false;
  }
  return true;
}

var instrumentJS = function (js) {
  if (program.analyzing === "false") return js;
  if (IsJsonString(js)) return js;
  const PREFIX = "window.__proxy__";
  const name = program.name;
  var addStack = true;
  var scriptNo = program.identifier;
  output = stateTracker.extractRelevantState(js, {
    PREFIX,
    name,
    addStack,
    provenance: false,
  });
  return output;
};

var removeIntegrityAttr = function (html) {
  var root = htmlparser.parse(html);
  var scripts = root.getElementsByTagName("script");
  for (var s of scripts) {
    s.removeAttribute("integrity");
  }
  return root.toString();
};

var instrumentHTML = function (html) {
  if (program.analyzing === "false") return html;
  html = removeIntegrityAttr(html);
  var dynLib = fs.readFileSync(DYNPATH, "utf8");
  return `<script>${dynLib}</script>` + html;
};

var main = function () {
  var input = fs.readFileSync(program.input, "utf8");
  var output;
  if (program.type.includes("javascript")) {
    output = instrumentJS(input);
  } else output = instrumentHTML(input);
  fs.writeFileSync(program.input, output);
};

main();
