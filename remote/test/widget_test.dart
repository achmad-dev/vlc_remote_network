// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility in the flutter_test package. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:remote/main.dart';

void main() {
  testWidgets('Counter increments smoke test', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(MyApp());

    // button icon prev and next
    expect(find.byIcon(Icons.skip_previous_sharp), findsOneWidget);
    expect(find.byIcon(Icons.skip_next_sharp), findsOneWidget);

    // button icon volume up and down
    expect(find.byIcon(Icons.arrow_upward_sharp), findsOneWidget);
    expect(find.byIcon(Icons.arrow_downward_sharp), findsOneWidget);

    // button icon play and pause
    expect(find.byIcon(Icons.pause_circle_outline_outlined), findsOneWidget);
  });
}
