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





L.control.layers(baseMaps).addTo(map);

// var crownHill = L.marker([39.75, -105.09]).bindPopup('This is Crown Hill Park.'),



// var openTopoMap = L.tileLayer('https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png', {
//     maxZoom: 19,
//     attribution: 'Map data: © OpenStreetMap contributors, SRTM | Map style: © OpenTopoMap (CC-BY-SA)'
// });

// layerControl.addBaseLayer(openTopoMap, "OpenTopoMap");

function addGeoJsonLayer(data) {
    // Create a feature group to hold all layers
    var allLayersGroup = L.featureGroup();

    for(var i = 0; i < data.length; i++){
        // Create the GeoJSON layer
        var geoJsonLayer = L.geoJSON(data[i].geom, {
            style: function(feature) {
                return {
                    color: 'blue',
                    fillColor: 'blue',
                    fillOpacity: 0.3
                };
            }
        });

        // Add each layer to the feature group
        allLayersGroup.addLayer(geoJsonLayer);
        
        // Add the GeoJSON layer to the map
        geoJsonLayer.addTo(map);
    }

    // Create overlay maps
    var overlayMaps = {
        "GeoJSON Layers": allLayersGroup
    };

    // Update layer control
    if (map.layerControl) {
        map.layerControl.remove();
    }
    map.layerControl = L.control.layers(baseMaps, overlayMaps).addTo(map);

    // Zoom to the bounds of all layers
    if (allLayersGroup.getLayers().length > 0) {
        map.fitBounds(allLayersGroup.getBounds(), {
            padding: [50, 50], // Optional: adds some padding
            maxZoom: 15 // Optional: prevents zooming in too close
        });
    }
}