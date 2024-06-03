import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/models/app/side_menu_data_model/side_menu_data.dart';
import 'package:project_sync/views/0.2_dashboard/dashboard.dart';
import 'package:project_sync/views/0.4_my_tasks/my_tasks.dart';
import 'package:project_sync/views/0.3_teams/teams.dart';
import 'package:project_sync/views/0.5_schedule/schedule.dart';
import 'package:project_sync/views/0.6_messages/messages.dart';
import 'package:project_sync/views/0.9_contacts/contacts.dart';

class LandingWidget extends StatelessWidget {
  const LandingWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<AnimateSideMenuBloc, double>(builder: (context, state) {
      return AnimatedPositioned(
        duration: const Duration(milliseconds: 500),
        top: 122,
        left: state,
        right: 0,
        bottom: 0,
        child: Builder(builder: (context) {
          final homeBloc = context.read<HomeBloc>();
          return PageView.builder(
            itemCount: SideMenuData.sideMenuData.length,
            controller: homeBloc.landingPageCtrl,
            itemBuilder: (context, index) {
              return Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 20, vertical: 15),
                child: returnBuildView(index),
              );
            },
          );
        }),
      );
    });
  }

  Widget returnBuildView(int index) {
    switch (index) {
      case 0:
        return const Dashboard();
      case 1:
        return const Teams();
      case 2:
        return const MyTasks();
      case 3:
        return const Schedule();
      case 4:
        return const Messages();
      case 5:
        return const Contacts();
      default:
        return const Dashboard();
    }
  }
}
