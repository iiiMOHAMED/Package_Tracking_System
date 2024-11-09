import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { orderCreationComponent } from './components/orderCreation/orderCreation.component';
import { ManageOrdersComponent } from './components/manage-orders/manage-orders.component';
const routes: Routes = [
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
  { path: 'orders', component: orderCreationComponent },
  { path: 'manage-orders', component: ManageOrdersComponent }, // New route for ManageOrdersComponent
  { path: '', redirectTo: '/register', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

/*{ path: 'orders/*',component: ManageOrdersComponent},*/
