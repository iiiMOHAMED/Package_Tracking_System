import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  name: string = '';
  email: string = '';
  phone: string = '';
  password: string = '';
  private apiUrl = 'http://localhost:8080/register';

  constructor(private http: HttpClient) {}

  onSubmit() {
    const user = {
      name: this.name,
      email: this.email,
      phone: this.phone,
      password: this.password,
    };

    this.http.post(this.apiUrl, user).subscribe({
      next: () => alert('Registration successful!'),
      error: (err) => alert('Registration failed: ' + err.error)
    });
  }
}
