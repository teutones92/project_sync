import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:window_manager/window_manager.dart';

import 'views/0.0_authentication/authentication.dart';
import 'views/0.1_home/home.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final pref = await SharedPreferences.getInstance();
  final logeIn = pref.getBool('isLoggedIn') ?? false;
  // pref.setBool("isLoggedIn", false);
  runApp(MainApp(logeIn: logeIn));
}

class MainApp extends StatelessWidget {
  const MainApp({super.key, required this.logeIn});
  final bool logeIn;

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: MultiBlocsProviderList.providers,
      child: BlocBuilder<DarkModeBloc, ThemeMode>(
        builder: (context, state) {
          return LayoutBuilder(
            builder: (_, cont) {
              if (cont.maxWidth <= 1200) {
                windowManager.setMinimumSize(const Size(1024, 800));
                context.read<AnimateSideMenuBloc>().isHide(true);
              } else {
                context.read<AnimateSideMenuBloc>().isHide(false);
              }
              return MaterialApp(
                debugShowCheckedModeBanner: false,
                themeMode: state,
                theme: ThemeX.lightTheme,
                darkTheme: ThemeX.darkTheme,
                home: logeIn ? const Home() : const Authentication(),
              );
            },
          );
        },
      ),
    );
  }
}
