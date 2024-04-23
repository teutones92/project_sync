import 'package:flutter/material.dart';

class CustomButton extends StatelessWidget {
  CustomButton({super.key, required this.text, required this.onPressed});
  final String text;
  final double textSize = 16;
  final VoidCallback onPressed;
  final double elevation = 5.0;
  final BorderRadius borderRadius = BorderRadius.circular(10);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: elevation,
      shape: RoundedRectangleBorder(
        borderRadius: borderRadius,
      ),
      color: Colors.transparent,
      child: Container(
        decoration: BoxDecoration(
          color: Colors.blue,
          borderRadius: borderRadius,
          gradient: const LinearGradient(
            colors: [Colors.blue, Colors.blueAccent],
            begin: Alignment.centerLeft,
            end: Alignment.centerRight,
          ),
        ),
        child: Expanded(
          child: InkWell(
            onTap: onPressed,
            child: Center(
                child: Text(
              text,
              style: TextStyle(fontSize: textSize),
            )),
          ),
        ),
      ),
    );
  }
}
