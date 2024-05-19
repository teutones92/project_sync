import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

extension NavigatorX on BuildContext {
  static Future<void> push(
      {required BuildContext context, required Widget page}) async {
    await Navigator.of(context).push(
      PageRouteBuilder(pageBuilder: (context, animation, secondaryAnimation) {
        return FadeTransition(
          opacity: animation,
          child: page,
        );
      }),
    );
  }

  static void pop(BuildContext context) => Navigator.of(context).pop();

  void pushReplacement(Widget page) {
    Navigator.of(this).pushReplacement(
      MaterialPageRoute(builder: (_) => page),
    );
  }

  static Future<void> pushAndRemoveUntil(
      {required BuildContext context, required Widget page}) async {
    await Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => page),
      (route) => false,
    );
  }
}

extension ThemeX on ThemeData {
  static TextStyle titleText =
      const TextStyle(fontSize: 30, fontWeight: FontWeight.bold);

  static get subtitleText =>
      const TextStyle(fontSize: 20, fontWeight: FontWeight.bold);

  static get errorTextStyle => const TextStyle(
        fontSize: 20,
        fontWeight: FontWeight.bold,
        color: Colors.red,
      );

  static ThemeData of(BuildContext context) => Theme.of(context);

  static const _goldColor = MaterialColor(0xFFFFD700, {
    50: Color(0xFFFFF2D4),
    100: Color(0xFFFFE7AA),
    200: Color(0xFFFFD700),
    300: Color(0xFFFFC300),
    400: Color(0xFFFFB800),
    500: Color(0xFFFFAB00),
    600: Color(0xFFFF9E00),
    700: Color(0xFFFF9300),
    800: Color(0xFFFF8700),
    900: Color(0xFFFF7600),
  });

  static Gradient get goldGradient => LinearGradient(
        colors: [
          _goldColor.shade200,
          _goldColor.shade800,
          _goldColor.shade900,
          _goldColor.shade800,
          _goldColor.shade200,
        ],
        begin: Alignment.centerLeft,
        end: Alignment.centerRight,
      );

  static const Color darkBlueColor = Color.fromARGB(255, 4, 1, 24);

  //** Light theme for the app*/

  static ThemeData get lightTheme => ThemeData(
        textTheme: GoogleFonts.lexendExaTextTheme().copyWith(
          bodyLarge: GoogleFonts.lexend().copyWith(
            color: Colors.black,
          ),
          bodyMedium: GoogleFonts.lexend().copyWith(
            color: Colors.black,
          ),
          bodySmall: GoogleFonts.lexend().copyWith(
            color: Colors.black,
          ),
        ),
        textButtonTheme: TextButtonThemeData(
          style: TextButton.styleFrom(
            foregroundColor: _goldColor.shade800,
          ),
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ElevatedButton.styleFrom(
            foregroundColor: _goldColor.shade800,
            backgroundColor: darkBlueColor,
            elevation: 5,
          ),
        ),
        buttonTheme: const ButtonThemeData(
          buttonColor: Colors.black,
          textTheme: ButtonTextTheme.primary,
        ),
        textSelectionTheme: TextSelectionThemeData(
          cursorColor: _goldColor.shade800,
          selectionColor: _goldColor.shade800,
          selectionHandleColor: _goldColor.shade800,
        ),
        cardColor: Colors.white,
        hoverColor: const Color.fromARGB(33, 194, 194, 194),
        focusColor: _goldColor.shade800,
        scaffoldBackgroundColor: Colors.grey.shade100,
        shadowColor: Colors.black,
        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.white,
          shadowColor: Colors.black,
          elevation: 5,
          iconTheme: IconThemeData(color: Colors.black),
        ),
        dialogTheme: DialogTheme(
          backgroundColor: Colors.white,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
        ),
        colorScheme: ColorScheme.fromSwatch(primarySwatch: _goldColor).copyWith(
          // surface: Colors.white,
          outline: _goldColor.shade800,
          brightness: Brightness.light,
        ),
      );

  //** Dark theme for the app*/

  static ThemeData get darkTheme => ThemeData(
        primarySwatch: _goldColor,
        textTheme: GoogleFonts.lexendExaTextTheme().copyWith(
          titleLarge: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          titleMedium: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          titleSmall: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          displayLarge: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          displayMedium: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          displaySmall: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          headlineLarge: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          headlineMedium: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          headlineSmall: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          bodyLarge: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          bodyMedium: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
          bodySmall: GoogleFonts.lexend().copyWith(
            color: Colors.white,
          ),
        ),
        listTileTheme: const ListTileThemeData(
          // tileColor: Colors.white,
          iconColor: Colors.white,
          textColor: Colors.white,
        ),
        textButtonTheme: TextButtonThemeData(
          style: TextButton.styleFrom(
            foregroundColor: _goldColor.shade800,
          ),
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ElevatedButton.styleFrom(
            foregroundColor: _goldColor.shade800,
            elevation: 5,
            disabledBackgroundColor: Colors.grey.shade800,
          ),
        ),
        iconButtonTheme: IconButtonThemeData(
          style: IconButton.styleFrom(
            foregroundColor: Colors.white,
            elevation: 5,
          ),
        ),
        textSelectionTheme: TextSelectionThemeData(
          cursorColor: _goldColor.shade800,
          selectionColor: _goldColor.shade800,
          selectionHandleColor: _goldColor.shade800,
        ),
        cardColor: darkBlueColor,
        scaffoldBackgroundColor: const Color.fromARGB(255, 10, 5, 0),
        dialogBackgroundColor: darkBlueColor,
        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.white,
          shadowColor: Colors.black,
          elevation: 5,
          iconTheme: IconThemeData(color: Colors.white),
          foregroundColor: Colors.white,
        ),
        colorScheme: ColorScheme.fromSwatch(primarySwatch: _goldColor).copyWith(
          surface: darkBlueColor,
          outline: _goldColor.shade800,
          brightness: Brightness.dark,
        ),
        hintColor: Colors.grey.shade700,
        hoverColor: const Color.fromARGB(33, 194, 194, 194),
        focusColor: _goldColor.shade800,
        buttonTheme: ButtonThemeData(
          buttonColor: _goldColor.shade800,
          textTheme: ButtonTextTheme.primary,
          focusColor: _goldColor.shade800,
        ),
      );
}
