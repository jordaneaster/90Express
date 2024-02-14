from PyQt5.QtWidgets import QLabel, QVBoxLayout, QWidget

class ProfilePage(QWidget):
    def __init__(self, username, email):
        super().__init__()

        # Create labels to display user information
        self.username_label = QLabel(f"Username: {username}")
        self.email_label = QLabel(f"Email: {email}")

        # Create layout and add labels
        layout = QVBoxLayout()
        layout.addWidget(self.username_label)
        layout.addWidget(self.email_label)

        self.setLayout(layout)
