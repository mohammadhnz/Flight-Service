# Generated by Django 4.1.1 on 2022-12-13 21:19

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ("bank", "0002_alter_transaction_callback"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="transaction",
            name="time_action",
        ),
    ]