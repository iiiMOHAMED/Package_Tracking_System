import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { orderCreationComponent } from './components/orderCreation/orderCreation.component';
import { ManageOrdersComponent } from './components/manage-orders/manage-orders.component';
import { UserOrdersComponent } from './user-orders/user-orders.component';
import { OrderDetailsComponent } from './order-details/order-details.component';
import { CourierComponent } from './courier/courier.component';
import { HomeComponent } from './home/home.component';
const routes: Routes = [
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
  { path: 'orders', component: orderCreationComponent },
  { path: 'manage-orders', component: ManageOrdersComponent }, // New route for ManageOrdersComponent
  { path: 'user-orders', component: UserOrdersComponent },
  { path: 'order-details/:orderNumber',component: OrderDetailsComponent},
  { path: 'courier',component: CourierComponent},
  { path: 'home',component: HomeComponent},
  { path: '', redirectTo: '/home', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

/*{ path: 'orders/*',component: ManageOrdersComponent},*/
