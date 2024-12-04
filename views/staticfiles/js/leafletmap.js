// Global variables
var map, osm, layerControl, geoJsonLayerGroup;

// Function to initialize the map and add custom input control
function initializeMap() {
    // Create map
    map = L.map('map', {
        center: [6.2, -1.99],
        zoom: 10,
        zoomControl: false // Disable default zoom control (optional)
    });

    // Add OpenStreetMap tile layer
    osm = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</>'
    }).addTo(map);

    // Create layer control with base maps
    layerControl = L.control.layers({
        "OpenStreetMap": osm
    }, {}).addTo(map);

    // Add zoom control explicitly (if it's disabled or removed)
    L.control.zoom({
        position: 'topleft' // Positioning it in the top-left
    }).addTo(map);
    


}

// Function to create the input element



function addGeoJsonLayer(data) {
    // Remove existing GeoJSON layers from the map and layer control
    if (geoJsonLayerGroup) {
        map.removeLayer(geoJsonLayerGroup);
        layerControl.removeLayer(geoJsonLayerGroup);
    }

    // Recreate the feature group
    geoJsonLayerGroup = L.featureGroup();

    // Add new GeoJSON layers
    for(var i = 0; i < data.length; i++){
        var geoJsonLayer = L.geoJSON(data[i].geom, {
            style: function(feature) {
                return {
                    color: 'blue',
                    fillColor: 'blue',
                    fillOpacity: 0.3
                };
            },
            onEachFeature: function(feature, layer) {
                // Create popup content with both shape__len and shape__area
                var popupContent = `
                    <div>
                        <strong>Shape Length:</strong> ${data[i].shape__len.toFixed(2)} meters<br>
                        <strong>Shape Area:</strong> ${data[i].shape__are.toFixed(2)} sq meters <br>
                        <strong>Is Storey?:</strong> ${data[i].building_t} <br>
                        <strong>building type:</strong> ${data[i].building_u} <br>
                        <strong>Creation day:</strong> ${data[i].creationda} <br>
                        <strong>Development Status:</strong> ${data[i].developmen} <br>
                        <strong>Exact Use:</strong> ${data[i].exact_use} <br>
                        <strong>Number of Story:</strong> ${data[i].num_storey} <br>
                        <strong>Plot Number:</strong> ${data[i].plot_numbe} <br>
                        <strong>Plot Number:</strong> ${data[i].remarks} <br>
                    </div>
                `;

                // Bind popup 
                layer.bindPopup(popupContent);

                // Optional: Add click event
                layer.on('click', function(e) {
                    console.log('Feature clicked:', {
                        length: data[i].shape__len,
                        area: data[i].shape__area
                    });
                });
            }
        });

        // Add layer to the feature group
        geoJsonLayerGroup.addLayer(geoJsonLayer);
    }

    // Add feature group to the map
    geoJsonLayerGroup.addTo(map);

    // Add to layer control
    layerControl.addOverlay(geoJsonLayerGroup, "GeoJSON Layers");

    // Zoom to layers if any exist
    if (geoJsonLayerGroup.getLayers().length > 0) {
        map.fitBounds(geoJsonLayerGroup.getBounds(), {
            padding: [50, 50],
            maxZoom: 15
        });
    }
}




initializeMap();

L.Control.SearchInput = L.Control.extend({
    options: {
        position: 'topleft' // Position of the control
    },

    onAdd: function (map) {
        const container = L.DomUtil.create('div', 'leaflet-control leaflet-control-search-input');

        const input = L.DomUtil.create('input', 'search-input', container);
        input.placeholder = 'Search here...';

        // Add event listener for keyup event to trigger on Enter key press
        input.addEventListener('keyup', function(event) {
            const searchTerm = event.target.value;

            // Check if the pressed key is Enter (key code 13)
            if (event.key === 'Enter') {
                console.log("Search term:", searchTerm);
                
                // Only send the request if there's a search term
                if (searchTerm.trim() !== '') {
                    // Example request
                    $.ajax({
                        url: '/map/simplesearch', // Your server-side URL
                        type: 'POST',
                        dataType: 'json', // Expecting JSON response
                        contentType: 'application/json',
                        data: JSON.stringify({ searchValue: searchTerm }),
                        success: function (response) {
                            if (response.success) {
                                console.log("Successfully received data:", response.data);
                                // Process response here (e.g., adding GeoJSON layers)
                                var data =  response.data
                                $(".results-returned").val(data.length)
                                addGeoJsonLayer(response.data);
                            } else {
                                alert("No results found.");
                            }
                        },
                        error: function (xhr, status, error) {
                            var errorMessage = xhr.responseJSON ? xhr.responseJSON.error : xhr.responseText || 'An unknown error occurred.';
                            alert("Error: " + errorMessage);
                        }
                    });
                } else {
                    alert("Please enter a valid search term.");
                }
            }
        });

        return container;
    }
});


// Add the control to the map
map.addControl(new L.Control.SearchInput());


L.Control.ResultsReturned = L.Control.extend({
    options: {
        position: 'bottomleft' // Position of the control
    },

    onAdd: function (map) {
        const container = L.DomUtil.create('div', 'leaflet-control leaflet-control-results-control-input');

        const input = L.DomUtil.create('input', 'results-returned', container);
        input.placeholder = 'no. of items appears here...';

        // Add event listener for search functionality
        input.addEventListener('keyup', function(event) {
            const searchTerm = event.target.value;

            
          
        });

        return container;
    }
});

// Add the control to the map
map.addControl(new L.Control.ResultsReturned());