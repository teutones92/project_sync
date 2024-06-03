import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';

class ProjectTeamMembers extends StatelessWidget {
  const ProjectTeamMembers({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Project Members',
          style: ThemeX.subtitleText,
        ),
        const SizedBox(height: 10),
        SizedBox(
          height: 50,
          child: Row(
            children: [
              Flexible(
                child: ListView.builder(
                  shrinkWrap: true,
                  itemCount: 5,
                  scrollDirection: Axis.horizontal,
                  itemBuilder: (context, index) {
                    return Padding(
                      padding: const EdgeInsets.only(right: 10),
                      child: Stack(
                        children: [
                          CircleAvatar(
                            radius: 30,
                            backgroundColor: Colors.grey.shade300,
                            child: const Icon(Icons.person),
                          ),
                          Positioned(
                            right: 0,
                            bottom: 0,
                            child: Container(
                              decoration: const BoxDecoration(
                                color: Colors.red,
                                shape: BoxShape.circle,
                              ),
                              child: IconButton(
                                padding: const EdgeInsets.all(2),
                                constraints: const BoxConstraints(),
                                onPressed: () {},
                                icon: const Icon(
                                  Icons.close,
                                  color: Colors.white,
                                  size: 15,
                                ),
                              ),
                            ),
                          ),
                        ],
                      ),
                    );
                  },
                ),
              ),
              Container(
                height: 50,
                width: 50,
                decoration: BoxDecoration(
                  color: Colors.grey.shade300,
                  shape: BoxShape.circle,
                ),
                child: IconButton(
                  onPressed: () {
                    // Show dialog to add new member
                    showDialog(
                        context: context,
                        builder: (context) => AlertDialog(
                              title: const Text('Select members'),
                              content: const Text(
                                  'Select members to add to the project'),
                              actions: [
                                TextButton(
                                  onPressed: () {
                                    Navigator.pop(context);
                                  },
                                  child: const Text('Cancel'),
                                ),
                                TextButton(
                                  onPressed: () {
                                    Navigator.pop(context);
                                  },
                                  child: const Text('Accept'),
                                ),
                              ],
                            ));
                  },
                  color: Colors.black,
                  icon: const Icon(Icons.add),
                ),
              )
            ],
          ),
        ),
      ],
    );
  }
}
