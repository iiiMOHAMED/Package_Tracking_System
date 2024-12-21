import { Component } from '@angular/core';
import { AuthService } from 'src/app/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  email: string = '';
  password: string = '';
  private apiUrl = 'https://backend2-crt-20216087-dev.apps.rm2.thpm.p1.openshiftapps.com/login';

  constructor(private authService: AuthService, private router: Router) {}

  onSubmit() {
    this.authService.login(this.email, this.password).subscribe({
      next: (response) => {
        // Check if response contains the token
        if (response && response.token) {
          this.authService.storeToken(response.token); // Store the token
          console.log("Token stored:", response.token);
          this.router.navigate(["/"]);
        } else {
          console.error("Login failed: Token not found in response");
        }
      },
      error: (err) => {
        console.error("Login failed", err);
      }
    });
  }
}
