import 'package:flutter/material.dart';

class ContactsBody extends StatelessWidget {
  const ContactsBody({super.key, required this.count, this.animate = false});
  final int count;
  final bool animate;

  @override
  Widget build(BuildContext context) {
    return GridView.builder(
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: count,
          crossAxisSpacing: 10,
          mainAxisSpacing: 10,
        ),
        itemBuilder: (context, index) {
          return Stack(
            children: [
              AnimatedPadding(
                duration: const Duration(milliseconds: 300),
                padding: EdgeInsets.all(animate ? 8 : 0),
                child: Card(
                  elevation: 8,
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            IconButton(
                                onPressed: () {},
                                icon: const Icon(Icons.star_border)),
                            IconButton(
                              onPressed: () {},
                              icon: const Icon(Icons.more_horiz),
                            ),
                          ],
                        ),
                        Column(
                          children: [
                            CircleAvatar(
                              radius: 30,
                              backgroundColor: Colors.grey.shade300,
                              child: const Icon(Icons.person),
                            ),
                            const ListTile(
                              title: Center(child: Text('Name')),
                              subtitle: Center(child: Text('Phone number')),
                            ),
                          ],
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            const Text('Aug 12, 2021'),
                            IconButton(
                              onPressed: () {},
                              icon: const Icon(Icons.message),
                            ),
                          ],
                        ),
                      ],
                    ),
                  ),
                ),
              ),
              Positioned(
                bottom: 0,
                right: 0,
                left: 0,
                child: AnimatedSwitcher(
                  duration: const Duration(milliseconds: 300),
                  child: animate
                      ? Container(
                          decoration: BoxDecoration(
                            color: Colors.green,
                            shape: BoxShape.circle,
                            boxShadow: [
                              BoxShadow(
                                color: Colors.black.withOpacity(0.2),
                                blurRadius: 5,
                                spreadRadius: 2,
                              ),
                            ],
                          ),
                          child: const Icon(Icons.check, size: 40),
                        )
                      : const SizedBox(),
                ),
              ),
            ],
          );
        });
  }
}
