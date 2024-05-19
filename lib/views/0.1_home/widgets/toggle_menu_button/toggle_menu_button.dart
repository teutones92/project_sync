import 'package:project_sync/core/config/bloc_config.dart';
import 'package:flutter/material.dart';

class ToggleMenuButton extends StatelessWidget {
  const ToggleMenuButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<AnimateSideMenuBloc, double>(
      builder: (context, state) {
        return AnimatedPositioned(
          duration: const Duration(milliseconds: 500),
          top: 40,
          left: state == 0 ? 30 : 328,
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).scaffoldBackgroundColor,
              shape: BoxShape.circle,
              border: Border.all(
                color: Theme.of(context).brightness == Brightness.light
                    ? Colors.grey.shade300
                    : Colors.grey.shade700,
                width: 1,
              ),
            ),
            child: Center(
              child: IconButton(
                onPressed: () {
                  context.read<AnimateSideMenuBloc>().animateSideMenu();
                },
                icon: Icon(
                  state == 0 ? Icons.menu_rounded : Icons.menu_open_rounded,
                  color: Theme.of(context).brightness == Brightness.light
                      ? Colors.grey.shade900
                      : Colors.grey.shade700,
                  weight: 0.1,
                ),
              ),
            ),
          ),
        );
      },
    );
  }
}
