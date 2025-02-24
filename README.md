Task: Student Management System with Goroutines

Objective:
Develop a Go program that manages student records and utilizes Goroutines to compute student grades concurrently.

Requirements:
1. The program should allow users to enter multiple student records, including:
   - First Name
   - Last Name
   - Math Score
   - Science Score
   - Literature Score

2. Implement a struct to store student data.
3. Use a Goroutine to calculate the average score for each student.
4. Display a list of all students with their scores and computed averages.
5. The program should be able to run calculations asynchronously.

Expected Output Format:
The program should display the list of students in the following format:

Student List:
1. John Doe | Math: 85 | Science: 90 | Literature: 88 | Average: 87.67
2. Alice Smith | Math: 92 | Science: 89 | Literature: 94 | Average: 91.67

Additional Challenge (Optional):
✅ Implement functionality to identify and display the top-performing student.
✅ Sort students in descending order based on their average scores.
✅ Save student data in a text file (CSV or JSON format) for future reference.

Note: As Goroutines execute concurrently, the output may not always appear in the same order.
