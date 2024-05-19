import 'package:flutter/material.dart';

import '../../core/extensions.dart';

class HeaderWidgetItems extends StatelessWidget {
  const HeaderWidgetItems({super.key, this.shadows = false});
  final bool shadows;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      alignment: WrapAlignment.start,
      crossAxisAlignment: WrapCrossAlignment.start,
      children: [
        SizedBox(
          width: 50,
          height: 50,
          child: Card(
            elevation: 5,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(10),
            ),
            color: ThemeX.darkBlueColor,
            child: Padding(
              padding: const EdgeInsets.only(top: 2),
              child: Center(
                child: Image.asset(
                  'assets/no_bg_icon.png',
                  width: 50,
                ),
              ),
            ),
          ),
        ),
        const SizedBox(width: 15),
        Text(
          "Project Sync.",
          style: ThemeX.titleText.copyWith(
            shadows: shadows
                ? [
                    const Shadow(
                      color: Colors.black,
                      blurRadius: 2,
                    ),
                  ]
                : null,
          ),
        ),
      ],
    );
  }
}
