from PyQt5.QtWidgets import QMainWindow, QVBoxLayout, QWidget, QLabel, QLineEdit, QPushButton, QMessageBox
from dashboard import Dashboard  # Import the Dashboard class
import requests

class LoginScreen(QMainWindow):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Login")
        self.setGeometry(100, 100, 400, 200)

        layout = QVBoxLayout()

        # UserID input
        self.username_label = QLabel("Username:")
        layout.addWidget(self.username_label)
        self.username_input = QLineEdit()
        layout.addWidget(self.username_input)

        # Password input
        self.password_label = QLabel("Password:")
        layout.addWidget(self.password_label)
        self.password_input = QLineEdit()
        self.password_input.setEchoMode(QLineEdit.Password)
        layout.addWidget(self.password_input)

        # Login button
        self.login_button = QPushButton("Login")
        self.login_button.clicked.connect(self.login)
        layout.addWidget(self.login_button)

        central_widget = QWidget()
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)

    def login(self):
        username = self.username_input.text()
        password = self.password_input.text()

        # Send a POST request to the authentication endpoint of your Go backend API
        url = "http://localhost:8080/api/authenticate"
        payload = {"username": username, "password": password}
        response = requests.post(url, json=payload)

        if response.status_code == 200:
            data = response.json()
            authenticated = data.get("authenticated", False)
            if authenticated:
                token = data.get("token")

                # Check if the user has created a profile
                profile_exists = self.check_profile_exists(username)

                if not profile_exists:
                    # If the profile doesn't exist, create a new profile
                    self.create_profile(username, token)

                QMessageBox.information(self, "Login", "Login successful!")
                self.dashboard = Dashboard(token,username)  # Create an instance of the Dashboard window
                self.dashboard.show()  # Show the Dashboard window
                self.close()  # Close the login screen
            else:
                QMessageBox.warning(self, "Login", "Invalid credentials. Please try again.")
        else:
            QMessageBox.warning(self, "Login", "Failed to connect to the server. Please try again later.")

    def check_profile_exists(self, username):
        # Send a GET request to your backend API to check if the user has already created a profile
        url = f"http://localhost:8080/api/check_profile?username={username}"
        response = requests.get(url)
        if response.status_code == 200:
            data = response.json()
            return data.get("profile_exists", False)
        else:
            QMessageBox.warning(self, "Error", "Failed to check profile existence.")
            return False

    def create_profile(self, username, token):
        # Send a POST request to your backend API to create a new profile for the user
        url = "http://localhost:8080/api/create_profile"
        payload = {"username": username}
        headers = {"Authorization": f"Bearer {token}"}
        response = requests.post(url, json=payload, headers=headers)
        if response.status_code == 200:
            QMessageBox.information(self, "Profile", "Profile created successfully.")
        else:
            QMessageBox.warning(self, "Error", "Failed to create profile.")

        # Send a POST request to the profile endpoint with the username in the request body
        profile_url = "http://localhost:8080/profile"
        profile_payload = {"username": username}
        profile_response = requests.post(profile_url, json=profile_payload, headers=headers)
        if profile_response.status_code != 200:
            QMessageBox.warning(self, "Error", "Failed to load profile page.")

