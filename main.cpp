#include <QApplication>
#include <QIcon>
#include <QMessageBox>
#include <QString>
#include <QStyle>
#include <QStyleFactory>
#include <QSystemTrayIcon>

#include "src/ai/L4Model.h"
#include "src/ui/MainWindow.h"

int
main(int argc, char* argv[]) {
  QApplication app(argc, argv);

  QApplication::setApplicationName("Dots and Boxes");
  QApplication::setApplicationVersion("1.0");
  QApplication::setOrganizationName("Dots and Boxes");
  QApplication::setStyle(QStyleFactory::create("Fusion"));

  L4Model model;
  auto mainWindow = std::make_unique<MainWindow>(true, true, &model, &model);
  mainWindow->show();

  return QApplication::exec();
}