import 'package:project_sync/core/config/bloc_config.dart';
import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/views/0.1_home/widgets/landing/landing.dart';
import 'package:project_sync/views/0.8_create_project/create_project.dart';
import 'widgets/home_header/home_header.dart';
import 'widgets/side_menu/side_menu.dart';
import 'widgets/toggle_menu_button/toggle_menu_button.dart';

class Home extends StatelessWidget {
  const Home({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          NavigatorX.push(context: context, page: const CreateProject());
        },
        child: const Icon(Icons.add),
      ),
      body: Stack(
        children: [
          BlocBuilder<AnimateSideMenuBloc, double>(
            builder: (context, state) {
              return AnimatedPositioned(
                duration: const Duration(milliseconds: 500),
                top: 120,
                left: state,
                child: Container(
                  width: MediaQuery.of(context).size.width,
                  height: 1,
                  decoration: BoxDecoration(
                    color: Theme.of(context).brightness == Brightness.light
                        ? Colors.grey.shade300
                        : Colors.grey.shade700,
                  ),
                ),
              );
            },
          ),
          const SideMenu(),
          const HomeHeader(),
          const ToggleMenuButton(),
          const LandingWidget(),
        ],
      ),
    );
  }
}
