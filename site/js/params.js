import {Window, HBox, VSlider, Button} from "./minicomps.min.mjs";

function loadParams(filepath) {
  return fetch(filepath)
    .then(response => response.json())
    .then(params => init(params));
}


function init(params) {
  const paramNames = Object.keys(params);
  const windowWidth = Math.max(paramNames.length * 45 + 12, 230);
  const win = new Window(document.body, 20, window.innerHeight - 290, windowWidth, 270, "Parameters");
  const box = new HBox(win, 20, 25, 30);
  paramNames.forEach(p => {
    new VSlider(box, 0, 0, p, params[p].value, params[p].min, params[p].max, (event) => {
      const slider = event.target;
      params[slider.label].min = slider.min;
      params[slider.label].max = slider.max;
      params[slider.label].value = slider.value;
      params[slider.label].decimals = slider.decimals * 10;
    }).setDecimals(params[p].decimals);
  });

  const buttonBox = new HBox(win, 10, 210, 10);
  new Button(buttonBox, 0, 0, "Send Params", () => {
    console.log(params);
  });
  new Button(buttonBox, 0, 0, "Re-compile");

}

export default loadParams;
