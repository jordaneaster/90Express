import sys
from PyQt5.QtWidgets import QApplication
from splash_screen import SplashScreen
from results_viewer import ResultsViewer

if __name__ == "__main__":
    app = QApplication(sys.argv)
    splash_screen = SplashScreen()
    splash_screen.show()
    sys.exit(app.exec_())
