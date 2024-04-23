import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

extension NavigatorX on BuildContext {
  void push(Widget page) {
    Navigator.of(this).push(
      MaterialPageRoute(builder: (_) => page),
    );
  }

  void pushReplacement(Widget page) {
    Navigator.of(this).pushReplacement(
      MaterialPageRoute(builder: (_) => page),
    );
  }

  void pushAndRemoveUntil(Widget page) {
    Navigator.of(this).pushAndRemoveUntil(
      MaterialPageRoute(builder: (_) => page),
      (route) => false,
    );
  }
}

extension ThemeX on ThemeData {
  Gradient get goldGradient => const LinearGradient(
        colors: [Color.fromRGBO(213, 186, 74, 1), Colors.blueAccent],
        begin: Alignment.centerLeft,
        end: Alignment.centerRight,
      );

  ThemeData get lightTheme => ThemeData(
        primarySwatch: Colors.blue,
        buttonTheme: const ButtonThemeData(
          buttonColor: Color.fromRGBO(213, 186, 74, 1),
          textTheme: ButtonTextTheme.primary,
        ),
      );

  ThemeData get darkTheme => ThemeData(
        primarySwatch: Colors.blue,
        scaffoldBackgroundColor: Colors.black,
        brightness: Brightness.dark,
        buttonTheme: const ButtonThemeData(
          buttonColor: Color.fromRGBO(213, 186, 74, 1),
          textTheme: ButtonTextTheme.primary,
        ),
      );
}
