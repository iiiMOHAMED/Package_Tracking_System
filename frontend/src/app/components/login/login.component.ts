import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
})
export class LoginComponent {
  email: string = '';
  password: string = '';
  private apiUrl = 'http://localhost:8080/login';

  constructor(private http: HttpClient) {}

  onSubmit() {
    const credentials = {
      email: this.email,
      password: this.password,
    };

    this.http.post(this.apiUrl, credentials).subscribe({
      next: () => alert('Login successful!'),
      error: (err) => alert('Login failed: ' + err.error)
    });
  }
}
