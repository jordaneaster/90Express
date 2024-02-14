from PyQt5.QtCore import QUrl
from PyQt5.QtWidgets import QMainWindow, QVBoxLayout, QWidget, QPushButton
from PyQt5.QtWebEngineWidgets import QWebEngineView

class Dashboard(QMainWindow):
    def __init__(self, token):
        super().__init__()
        self.setWindowTitle("Dashboard")
        self.setGeometry(100, 100, 800, 600)

        layout = QVBoxLayout()

        self.web_view = QWebEngineView()
        self.web_view.load(QUrl("http://localhost:8080/dashboard"))
        layout.addWidget(self.web_view)

        central_widget = QWidget()
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)

        self.web_view.loadFinished.connect(self.on_load_finished)

    def on_load_finished(self):
        self.web_view.page().runJavaScript('''
            var profileButton = document.querySelector("#profile-button");
            if (profileButton) {
                profileButton.click();
            }
        ''')
        
