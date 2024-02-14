import requests
import webbrowser
import folium
import os

class Map:
    def __init__(self, center, zoom_start):
        self.center = center
        self.zoom_start = zoom_start

    def showMap(self, features):
        # Create the map
        my_map = folium.Map(location=self.center, zoom_start=self.zoom_start)

        # Add markers based on the features obtained from the API
        for feature in features:
            for polygon, points in feature['features'].items():
                for point in points:
                    folium.Marker(location=[point[1], point[0]], popup=f"Inside {polygon}", tooltip='Click here to see Popup').add_to(my_map)

        # Save the map to an HTML file and open it in the default web browser
        current_dir = os.path.dirname(os.path.abspath(__file__))
        map_path = os.path.join(current_dir, 'map.html')
        my_map.save(map_path)
        webbrowser.get('safari').open(map_path)


if __name__ == "__main__":
    # Define the center and zoom level of the map
    center = [37.6000, -95.6650]
    zoom_start = 4.25

    # Fetch data from the Go backend API
    response = requests.get("http://127.0.0.1:8080/api/process")
    if response.status_code == 200:
        features = response.json()
        # Create the map and display it
        my_map = Map(center=center, zoom_start=zoom_start)
        my_map.showMap(features)
    else:
        print("Error fetching data from the API")
