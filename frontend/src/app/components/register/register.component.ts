import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  name: string = '';
  email: string = '';
  phone: string = '';
  password: string = '';
  role: string='';
  private apiUrl = 'http://localhost:8080/register';

  constructor(private http: HttpClient, private router: Router) {}

  onSubmit() {
    const user = {
      name: this.name,
      email: this.email,
      phone: this.phone,
      password: this.password,
      role: this.role
    };

    this.http.post(this.apiUrl, user).subscribe({
      next: () =>{ alert('Registration successful!');
                   this.router.navigate(["/"]);
      },
      error: (err) => alert('Registration failed: ' + err.error)
    });
  }
}
