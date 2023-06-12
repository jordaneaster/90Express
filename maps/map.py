import webbrowser
import folium

if __name__ == "__main__":

    class Map:

        def __init__(self, center, zoom_start):
            self.center = center
            self.zoom_start = zoom_start

        def showMap(self):
                fig3=folium.Figure(width=550,height=350)
                #Create the map
                my_map = folium.Map(location = self.center, zoom_start = self.zoom_start)
                fig3.add_child(my_map)
                #Define markers
                #Buffalo
                folium.Marker(location=[42.8864, 78.8784],popup='Default popup Marker1',tooltip='Click here to see Popup').add_to(my_map)
                #Erie
                folium.Marker(location=[28.695800, 80.0851],popup='Default popup Marker1',tooltip='Click here to see Popup').add_to(my_map)
                #Cleveland
                folium.Marker(location=[41.4993, 81.6944],popup='Default popup Marker1',tooltip='Click here to see Popup').add_to(my_map)

                #Display the map
                my_map.save("map.html")
                webbrowser.open("map.html")
        
    coords=[37.6000,-95.6650]
    map = Map(center = coords, zoom_start = 4.25)
    map.showMap()
