from PyQt5.QtCore import QUrl, QObject, pyqtSignal, pyqtSlot, QUrlQuery
from PyQt5.QtCore import QUrl
from PyQt5.QtWidgets import QMainWindow, QVBoxLayout, QWidget
from PyQt5.QtWebEngineWidgets import QWebEngineView
from PyQt5.QtWebChannel import QWebChannel
from PyQt5.QtCore import QObject, pyqtSignal, pyqtSlot

class Bridge(QObject):
    showProfileForm = pyqtSignal(str)  # Modify the signal to accept a string parameter

    @pyqtSlot()
    def displayProfileForm(self, username):  # Modify the method signature to accept username
        self.showProfileForm.emit(username)
        
class Dashboard(QMainWindow):
    def __init__(self, token, username):
        super().__init__()
        self.setWindowTitle("Dashboard")
        self.setGeometry(100, 100, 800, 600)

        layout = QVBoxLayout()

        self.web_view = QWebEngineView()
        layout.addWidget(self.web_view)

        central_widget = QWidget()
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)

        self.bridge = Bridge()
        self.bridge.showProfileForm.connect(self.display_profile_form)

        channel = QWebChannel()
        channel.registerObject("bridge", self.bridge)
        self.web_view.page().setWebChannel(channel)

        url = QUrl("http://localhost:8080/profile")
        query = QUrlQuery()
        query.addQueryItem("username", username)
        url.setQuery(query.toString())

        self.web_view.load(url)

    def display_profile_form(self, username):
        # Construct the URL for the profile page with the username as a query parameter
        url = QUrl("http://localhost:8080/profile")
        query = QUrlQuery()
        query.addQueryItem("username", username)
        url.setQuery(query.toString())
        self.web_view.load(url)

    def reset_window(self):
        self.web_view.setUrl(QUrl())  # Clear the content
        self.load_dashboard()  # Reload the dashboard
