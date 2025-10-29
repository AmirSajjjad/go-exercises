# Go Exercises

A collection of small Go exercises and mini projects designed to help you practice and improve your Go programming skills.  
Each exercise lives in its own folder with source code and a dedicated `README.md` explaining the problem, concepts, and key takeaways.

---

## 📁 Repository Structure

```
go-exercises/
├── README.md            # This file
├── .gitignore
├── LICENSE
├── exercises/           # All exercises and mini projects
│   ├── 01-concurancy-read-file/
│   │   ├── read_file.go
│   │   ├── go.mod
│   │   └── README.md
│   ├── ...              # Comming Soon
│   └── ...              # Comming Soon
└── docs/                # General Go notes or tips (Comming Soon)
```

---

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/AmirSajjjad/go-exercises.git
   cd go-exercises
   ```

2. Navigate to an exercise folder and run the code:
   ```bash
   cd exercises/01-hello-world
   go run main.go
   ```

3. Each exercise has its own `README.md` file that explains:
   - The goal of the exercise
   - Key Go concepts
   - Example input/output
   - Additional notes or learning tips

---

## 🧠 Tips
- Name folders in a numbered format, e.g., `01-hello-world`, `02-slices-maps`, etc.
- Always initialize Go modules inside each exercise if needed:
  ```bash
  go mod init github.com/<your-username>/go-exercises/exercises/01-hello-world
  ```
- Add tests using `_test.go` files to practice unit testing.

---

## 🧩 Contributing
If you'd like to contribute new exercises or improvements, feel free to fork the repo and open a pull request.

---

## 📜 License
This project is open-source under the [MIT License](LICENSE).

Happy coding with Go! 🐹
