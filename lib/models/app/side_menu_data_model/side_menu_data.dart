import 'package:flutter/material.dart';

class SideMenuData {
  final String title;
  final IconData icon;
  final bool selected;

  SideMenuData({
    required this.title,
    required this.icon,
    required this.selected,
  });

  static final List<SideMenuData> sideMenuData = [
    SideMenuData(
      title: "Dashboard",
      icon: Icons.dashboard,
      selected: true,
    ),
    SideMenuData(
      title: "Teams",
      icon: Icons.group,
      selected: false,
    ),
    SideMenuData(
      title: "My Tasks",
      icon: Icons.task,
      selected: false,
    ),
    SideMenuData(
      title: "Schedule",
      icon: Icons.calendar_today,
      selected: false,
    ),
    SideMenuData(
      title: "Messages",
      icon: Icons.message,
      selected: false,
    ),
    SideMenuData(
      title: "Settings",
      icon: Icons.settings,
      selected: false,
    ),
  ];
}
