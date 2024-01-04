napConsoleLog = function () {
  console.log(...arguments)
}


function napInit(devMode) {
  console.log("%cNAP INIT", "font-size: 1.5em; color: lightblue; font: courier; font-weight: bold")
  napReleaseObserver.observe(document.documentElement, napReleaseConfig);
  if (devMode) {
    window.addEventListener("click", napExitClick)
  }
}

function napExitClick(event) {
  // allows ctrl + mouse left click to close the chrome session gracefully
  if (event.shiftKey && !event.ctrlKey && !event.altKey) {
    let chromeDpDone = document.createElement("div");
    chromeDpDone.setAttribute("id", "chromeDpDone")
    document.documentElement.append(chromeDpDone)
  }
}

const napReleaseConfig = {childList: true, subtree: true};

const napReleaseObserver = new MutationObserver(function (mutations, observer) {
  mutations.forEach((value) => value.removedNodes.forEach(napReleaseNode))
})

const ReleaseFunc = "napRelease"

napReleaseId = function (id) {
  console.log("napReleaseId not overridden", id)
}

function napReleaseNode(node) {
  node.childNodes.forEach(napReleaseNode)
  if ("data-nap-id" in node) {
    // console.log("napReleaseNode", node["data-nap-id"] )
    napReleaseId(node["data-nap-id"])
  }
  // if (ReleaseFunc in node) {
  //   node[ReleaseFunc]()
  // }
}

