import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:flutter/material.dart';

class IconDarkMode extends StatelessWidget {
  const IconDarkMode({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<DarkModeBloc, ThemeMode>(
      builder: (context, state) {
        return IconButton(
          onPressed: () {
            context.read<DarkModeBloc>().toggleDarkMode();
          },
          icon: AnimatedSwitcher(
            duration: const Duration(milliseconds: 300),
            child: Icon(
              key: ValueKey(state),
              state == ThemeMode.dark ? Icons.dark_mode : Icons.light_mode,
              color: state == ThemeMode.dark
                  ? Colors.grey.shade200
                  : ThemeX.goldGradient.colors[0],
            ),
          ),
        );
      },
    );
  }
}
