import 'package:flutter/material.dart';

class NotificationButton extends StatelessWidget {
  const NotificationButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.grey.withOpacity(0.1),
        shape: BoxShape.circle,
      ),
      padding: const EdgeInsets.all(3),
      child: Center(
        child: IconButton(
          onPressed: () {},
          icon: const Stack(
            children: [
              Positioned(
                right: 4,
                child: Badge(
                  backgroundColor: Color.fromARGB(255, 255, 17, 0),
                ),
              ),
              Icon(Icons.notifications_none_rounded),
            ],
          ),
        ),
      ),
    );
  }
}
