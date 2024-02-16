import sys
from PyQt5.QtWidgets import QApplication
from splash_screen import SplashScreen
from PyQt5.QtWebEngineWidgets import QWebEngineSettings


if __name__ == "__main__":
    app = QApplication(sys.argv)
    splash_screen = SplashScreen()
    splash_screen.show()
    sys.exit(app.exec_())
# In your QMainWindow or wherever appropriate
settings = QWebEngineSettings.globalSettings()
settings.setAttribute(QWebEngineSettings.LocalStorageEnabled, False)
settings.setAttribute(QWebEngineSettings.LocalContentCanAccessRemoteUrls, True)
settings.setAttribute(QWebEngineSettings.LocalContentCanAccessFileUrls, True)
settings.setAttribute(QWebEngineSettings.AllowRunningInsecureContent, True)
settings.setAttribute(QWebEngineSettings.JavascriptCanAccessClipboard, True)
settings.setAttribute(QWebEngineSettings.PluginsEnabled, True)
settings.setAttribute(QWebEngineSettings.ScrollAnimatorEnabled, True)
settings.setAttribute(QWebEngineSettings.FullScreenSupportEnabled, True)
settings.setAttribute(QWebEngineSettings.ScreenCaptureEnabled, True)
settings.setAttribute(QWebEngineSettings.ErrorPageEnabled, True)
settings.setAttribute(QWebEngineSettings.AutoLoadImages, True)
settings.setAttribute(QWebEngineSettings.JavascriptEnabled, True)
settings.setAttribute(QWebEngineSettings.JavascriptCanOpenWindows, True)
settings.setAttribute(QWebEngineSettings.JavascriptCanAccessClipboard, True)
settings.setAttribute(QWebEngineSettings.LocalContentCanAccessRemoteUrls, True)
settings.setAttribute(QWebEngineSettings.LocalContentCanAccessFileUrls, True)
settings.setAttribute(QWebEngineSettings.PluginsEnabled, True)
settings.setAttribute(QWebEngineSettings.ScreenCaptureEnabled, True)
settings.setAttribute(QWebEngineSettings.ErrorPageEnabled, True)
settings.setAttribute(QWebEngineSettings.AutoLoadImages, True)