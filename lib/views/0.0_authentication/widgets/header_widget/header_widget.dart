import 'package:project_sync/global/widgets/header_widget_items.dart';
import 'package:project_sync/utils/icon_dark_mode.dart';
import 'package:flutter/material.dart';

class HeaderWidget extends StatelessWidget {
  const HeaderWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return const Stack(
      children: [
        Positioned(
          left: 60,
          top: 30,
          right: 60,
          child: Wrap(
            alignment: WrapAlignment.spaceBetween,
            crossAxisAlignment: WrapCrossAlignment.center,
            children: [
              HeaderWidgetItems(),
              IconDarkMode(),
            ],
          ),
        ),
      ],
    );
  }
}
