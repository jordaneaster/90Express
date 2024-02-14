from PyQt5.QtWidgets import QMainWindow, QVBoxLayout, QWidget, QLabel, QLineEdit, QPushButton, QMessageBox, QTextEdit
from PyQt5.QtCore import Qt
import requests

class ResultsViewer(QMainWindow):
    def __init__(self, token):
        super().__init__()
        self.setWindowTitle("Results Viewer")
        self.token = token
        
        self.text_area = QTextEdit()
        self.load_button = QPushButton("Load Results")
        self.load_button.clicked.connect(self.load_results)

        central_widget = QWidget()
        layout = QVBoxLayout()
        layout.addWidget(self.text_area)
        layout.addWidget(self.load_button)
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)

    def load_results(self):
        headers = {"Authorization": f"Bearer {self.token}"}
        # Make a request to your Go backend API
        response = requests.get("http://localhost:8080/api/process", headers=headers)
        if response.status_code == 200:
            features = response.json()
            # Display the results in the text area
            self.text_area.clear()
            for feature in features:
                for polygon, points in feature['features'].items():
                    self.text_area.insertPlainText(f"Polygon: {polygon}\n")
                    for point in points:
                        self.text_area.insertPlainText(f"Latitude: {point[1]}, Longitude: {point[0]}\n")
                    self.text_area.insertPlainText("\n")
        else:
            self.text_area.clear()
            self.text_area.insertPlainText("Error fetching data from the API\n")
