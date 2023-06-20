import folium
#python library
import pandas

#pandas structures this data by lables at header of txt file
vol=pandas.read_csv('90xdatabase.txt')
#vol objects are the lables inside the txt file
name=list(vol['city'])
st=list(vol['state'])
ln=list(vol['log'])
lt=list(vol['lat'])

#functions that colors the dots on map can add conditions for differnt colors based on say location size etc...
def color_pings():
     return 'blue'

#start location of map and style of the basemap
map=folium.Map(location=[41.4993,-81.6944],zoom_start=6,tiles='Stamen Terrain')

#can add layers of pins that can be turned on & off that could show say differnt companys and also remove the 'stamen terrain' styling on basemap
fg=folium.FeatureGroup(name='90 express hubs')

#sets up markers location if they have a feature like "pop up" and the style/shape of markers
for n,st,ln,lt in zip(name,st,ln,lt):
 fg.add_child(folium.CircleMarker(location=[lt,ln],popup=str(n)+':'+str(st),fill_color=color_pings(),color=color_pings(),fill_opacity=0.7))


#the builder functions 
map.add_child(fg)
map.add_child(folium.LayerControl())
map.save('90xMap1.html')
map.show_in_browser()