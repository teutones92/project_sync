import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/global/widgets/header_widget_items.dart';
import 'package:flutter/material.dart';

import 'widgets/menu_buttons_widget.dart';

class SideMenu extends StatelessWidget {
  const SideMenu({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<AnimateSideMenuBloc, double>(
      builder: (context, state) {
        return AnimatedPositioned(
          top: 0,
          left: state == 0 ? -350 : 0,
          bottom: 0,
          duration: const Duration(milliseconds: 500),
          width: 350,
          child: Container(
            decoration: BoxDecoration(
              border: Border(
                right: BorderSide(
                  color: Theme.of(context).brightness == Brightness.light
                      ? Colors.grey.shade300
                      : Colors.grey.shade700,
                  width: 1,
                ),
              ),
            ),
            child: Padding(
              padding: const EdgeInsets.all(20),
              child: ListView(
                children: const [
                  HeaderWidgetItems(shadows: true),
                  SizedBox(height: 70),
                  MenuButtonsWidget(),
                  SizedBox(height: 10),
                  LogOutButton(),
                  SizedBox(height: 10),
                ],
              ),
            ),
          ),
        );
      },
    );
  }
}

class LogOutButton extends StatelessWidget {
  const LogOutButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: const Icon(
        Icons.logout_rounded,
        color: Colors.grey,
      ),
      onTap: () => showDialog(
        context: context,
        builder: (_) => AlertDialog(
          title: const Text("Logout"),
          content: const Text("Are you sure you want to logout?"),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text("Cancel"),
            ),
            TextButton(
              onPressed: () => context.read<HomeBloc>().logOut(context),
              child: const Text("Logout"),
            ),
          ],
        ),
      ),
      title: const Text(
        "Logout",
        style: TextStyle(fontWeight: FontWeight.bold, color: Colors.grey),
      ),
    );
  }
}
