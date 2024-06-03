import 'package:flutter/material.dart';

class SideMenuData {
  final String title;
  final IconData icon;

  SideMenuData({
    required this.title,
    required this.icon,
  });

  static final List<SideMenuData> sideMenuData = [
    SideMenuData(
      title: "Dashboard",
      icon: Icons.dashboard,
    ),
    SideMenuData(
      title: "Teams",
      icon: Icons.group,
    ),
    SideMenuData(
      title: "My Tasks",
      icon: Icons.task,
    ),
    SideMenuData(
      title: "Schedule",
      icon: Icons.calendar_today,
    ),
    SideMenuData(
      title: "Messages",
      icon: Icons.message,
    ),
    SideMenuData(
      title: "Contacts",
      icon: Icons.contact_phone,
    ),
    SideMenuData(
      title: "Settings",
      icon: Icons.settings,
    ),
  ];
}
