import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/models/app/side_menu_data_model/side_menu_data.dart';

class MenuButtonsWidget extends StatelessWidget {
  const MenuButtonsWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Builder(builder: (context) {
      final sideMenuBloc = context.read<SideMenuBloc>();
      return Column(
        children: List.generate(SideMenuData.sideMenuData.length, (index) {
          final data = SideMenuData.sideMenuData[index];
          return Padding(
            padding: const EdgeInsets.symmetric(vertical: 10),
            child: Card(
              elevation: data.selected ? 5 : 0,
              color: data.selected
                  ? ThemeX.goldGradient.colors[1]
                  : Colors.transparent,
              child: ListTile(
                leading: Icon(
                  data.icon,
                  color: data.selected ? Colors.white : Colors.grey,
                ),
                title: Text(
                  data.title,
                  style: TextStyle(
                      color: data.selected ? Colors.white : Colors.grey,
                      fontWeight: FontWeight.bold),
                ),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
                onTap: () => sideMenuBloc.onTap(index: index, context: context),
              ),
            ),
          );
        }),
      );
    });
  }
}
