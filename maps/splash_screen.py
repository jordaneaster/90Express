from PyQt5.QtWidgets import QMainWindow, QLabel, QVBoxLayout, QPushButton, QWidget
from PyQt5.QtGui import QPixmap
from PyQt5.QtCore import Qt
from login_screen import LoginScreen

class SplashScreen(QMainWindow):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("90Express ZoneMapDB")
        self.setGeometry(100, 100, 500, 300)

        layout = QVBoxLayout()

        # Add logo
        logo_label = QLabel()
        pixmap = QPixmap("90x.jpg")  # Replace "logo.png" with the path to your logo image file
        pixmap = pixmap.scaled(600,500)
        logo_label.setPixmap(pixmap)
        logo_label.setAlignment(Qt.AlignCenter)
        layout.addWidget(logo_label)

        # Add animation
        # animation_label = QLabel("Animating Splash Screen...")
        # animation_label.setAlignment(Qt.AlignCenter)
        # layout.addWidget(animation_label)

        # Add start button
        start_button = QPushButton("Get Started")
        start_button.clicked.connect(self.start_application)
        layout.addWidget(start_button)

        central_widget = QWidget()
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)

        # Perform animation
        # self.animation = QPropertyAnimation(animation_label, b"geometry")
        # self.animation.setDuration(1000)
        # self.animation.setStartValue(animation_label.geometry())
        # self.animation.setEndValue(animation_label.geometry().translated(0, 100))
        # self.animation.setLoopCount(-1)
        # self.animation.start()

    def start_application(self):
        # Open the LoginScreen window
        self.login_screen = LoginScreen()
        self.login_screen.show()
        self.close()
