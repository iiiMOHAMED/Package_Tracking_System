import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <h1>Welcome to the Shipping App</h1>
    <nav>
      <a routerLink="/register">Register</a> |
      <a routerLink="/login">Login</a> |
      <a routerLink="/orders">Order</a> |
      <a routerLink="/manage-orders">Manage Orders</a> <!-- New link for Manage Orders -->
    </nav>
    <router-outlet></router-outlet>
  `,
  styles: [
    `nav { margin: 10px 0; }
     a { margin-right: 10px; text-decoration: none; }
    `
  ]
})
export class AppComponent {
  title = 'shipping-frontend';
}
