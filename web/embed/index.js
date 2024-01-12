import * as React from "react";
import ReactDOMServer from "react-dom/server";
import App from "../src/App";

function renderToStringEmbed() {
  return ReactDOMServer.renderToString(<App />);
}

global.__embed = renderToStringEmbed;
