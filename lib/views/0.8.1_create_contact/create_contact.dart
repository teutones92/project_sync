import 'package:flutter/material.dart';

class CreateContact extends StatelessWidget {
  const CreateContact({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: CustomScrollView(
      slivers: [
        const SliverAppBar(
          title: Text('Create Contact'),
          floating: true,
          snap: true,
        ),
        SliverList(
          delegate: SliverChildListDelegate(
            [
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: Column(
                  children: [
                    const TextField(
                      decoration: InputDecoration(labelText: 'Name'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Email'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Phone'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Address'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'City'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'State'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Zip'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Country'),
                    ),
                    const TextField(
                      decoration: InputDecoration(labelText: 'Notes'),
                    ),
                    ElevatedButton(
                      onPressed: () {},
                      child: const Text('Save'),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ],
    ));
  }
}
