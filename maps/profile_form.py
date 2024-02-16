# profile_form.py

from PyQt5.QtWidgets import QDialog, QVBoxLayout, QLabel, QLineEdit, QPushButton

class ProfileForm(QDialog):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Create Profile")
        
        layout = QVBoxLayout()

        self.username_edit = QLineEdit()
        layout.addWidget(QLabel("Username:"))
        layout.addWidget(self.username_edit)

        self.email_edit = QLineEdit()
        layout.addWidget(QLabel("Email:"))
        layout.addWidget(self.email_edit)

        self.fullname_edit = QLineEdit()
        layout.addWidget(QLabel("Full Name:"))
        layout.addWidget(self.fullname_edit)

        self.company_edit = QLineEdit()
        layout.addWidget(QLabel("Company:"))
        layout.addWidget(self.company_edit)

        submit_button = QPushButton("Submit")
        submit_button.clicked.connect(self.submit_profile)
        layout.addWidget(submit_button)

        self.setLayout(layout)

    def submit_profile(self):
        # Retrieve data from form fields and submit the profile
        username = self.username_edit.text()
        email = self.email_edit.text()
        fullname = self.fullname_edit.text()
        company = self.company_edit.text()

        # Send the profile data to the server for processing (implement this part)

        # Close the dialog after submitting the profile
        self.accept()
