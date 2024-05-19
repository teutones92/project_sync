import 'package:flutter/material.dart';

import '../core/config/bloc_config.dart';

class DarkModeBloc extends Cubit<ThemeMode> {
  DarkModeBloc() : super(ThemeMode.system);

  void toggleDarkMode() {
    if (state == ThemeMode.light) {
      emit(ThemeMode.dark);
    } else {
      emit(ThemeMode.light);
    }
  }

  void setDarkMode(bool isDark) {
    emit(isDark ? ThemeMode.dark : ThemeMode.light);
  }

  void setSystemMode() {
    emit(ThemeMode.system);
  }
}
