def get_valid_amount(prompt):
    while True:
        try:
            amount = float(input(prompt))
            if amount <= 0:
                print("Amount must be greater than 0.")
                continue
            return amount
        except ValueError:
            print("Invalid input. Please enter a valid number.")


def main():
    balance = 0.0
    history = []

    print("========================================")
    print("         WELCOME TO RUTH'S BANK         ")
    print("========================================")
    print("This banking app lets you manage your money securely with love.")
    print("Choose an option from the menu below.")
    print("========================================")

    while True:
        print("\nMAIN MENU")
        print("1. Deposit money")
        print("2. Withdraw money")
        print("3. Check balance")
        print("4. View transaction history")
        print("5. Exit")

        try:
            choice = int(input("Enter your choice (1-5): "))
        except ValueError:
            print("Please enter a valid number.")
            continue

        if choice == 1:
            amount = get_valid_amount("Enter amount to deposit: ")
            balance += amount
            history.append(f"Deposited ${amount:.2f}")
            print(f"Deposit successful. New balance: ${balance:.2f}")

        elif choice == 2:
            amount = get_valid_amount("Enter amount to withdraw: ")
            if amount > balance:
                print("Insufficient funds. Your current balance is too low.")
            else:
                balance -= amount
                history.append(f"Withdrew ${amount:.2f}")
                print(f"Withdrawal successful. New balance: ${balance:.2f}")

        elif choice == 3:
            print(f"Your current balance is: ${balance:.2f}")

        elif choice == 4:
            if not history:
                print("No transactions yet.")
            else:
                print("\nTransaction History:")
                for item in history:
                    print("-", item)

        elif choice == 5:
            print("\n========================================")
            print("Thank you for using SMART BANK.")
            print("Your final balance is: $%.2f" % balance)
            if history:
                print("You completed", len(history), "transaction(s).")
            else:
                print("You completed 0 transactions.")
            print("Have a great day! Goodbye.")
            print("========================================")
            break

        else:
            print("Invalid choice. Please select a number from 1 to 5.")


if __name__ == "__main__":
    main()
