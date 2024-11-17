var osm = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</>'
});


var map = L.map('map', {
    center: [6.2, -1.99],
    zoom: 10,
    layers: [osm]
});

var baseMaps = {
    "OpenStreetMap": osm,
};


// var overlayMaps = {
//     "Cities": cities
// };

var layerControl = L.control.layers(baseMaps).addTo(map);

// var crownHill = L.marker([39.75, -105.09]).bindPopup('This is Crown Hill Park.'),



var openTopoMap = L.tileLayer('https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: 'Map data: © OpenStreetMap contributors, SRTM | Map style: © OpenTopoMap (CC-BY-SA)'
});

layerControl.addBaseLayer(openTopoMap, "OpenTopoMap");
layerControl.addOverlay(parks, "Parks");

