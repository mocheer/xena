var d3array  = require('./d3-array')
var d3contour = require('./d3-contour')
var data = require('./grid.json')

console.time('contours')
d3.contours().size([687,412]).thresholds([-9999,0.5, 1, 2])(data.Data.flat())
console.timeEnd('contours')
