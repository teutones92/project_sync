import 'package:flutter/material.dart';

class DashboardAnalyticModel {
  final String subTitle;
  final int value;
  final IconData icon;
  final Color color;

  DashboardAnalyticModel({
    required this.subTitle,
    required this.value,
    required this.icon,
    required this.color,
  });

  static final List<DashboardAnalyticModel> data = [
    DashboardAnalyticModel(
      subTitle: 'In Progress Projects',
      value: 0,
      icon: Icons.edit_calendar_outlined,
      color: Colors.orange,
    ),
    DashboardAnalyticModel(
      subTitle: 'Completed Projects',
      value: 0,
      icon: Icons.check_box_rounded,
      color: Colors.green,
    ),
    DashboardAnalyticModel(
      subTitle: 'Today\'s Tasks',
      value: 0,
      icon: Icons.task_rounded,
      color: Colors.blue,
    ),
    DashboardAnalyticModel(
      subTitle: 'Total Members',
      value: 0,
      icon: Icons.people_alt_rounded,
      color: Colors.purple,
    ),
  ];
}
