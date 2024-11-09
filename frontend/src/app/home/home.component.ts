import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { jwtDecode } from "jwt-decode";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  userId:number|null=null;
  role:string|null=null;
  constructor(private r:Router,private authService:AuthService,){
    const token = this.authService.getToken();
    if (token) {
      const decodedToken: any = jwtDecode(token);
      this.userId = decodedToken.sub; 
      this.role=decodedToken.role;
      console.log("test1:"+this.userId)
      console.log("test1:"+this.role)
    }
    else{ 
      console.log("test2:"+this.userId)
      console.log("test2:"+this.role)
    }
  }


  ngOnInit(): void {
    if (!this.isLoggedIn()) {
      this.r.navigate(['/login']);  // Redirect if not logged in
    }
  }
  goLogin(){
    this.r.navigate(["/login"])
  }
  goRegister(){
    this.r.navigate(["/register"])
  }
  goCreateOrder(){
    this.r.navigate(["/orders"])
  }
  goManageOrders(){
    this.r.navigate(["/manage-orders"])
  }
  goMyOrders(){
    this.r.navigate(["/user-orders"])
  }
  goAssignedOrders(){
    this.r.navigate(["/courier"])
  }

  isLoggedIn(){
    return this.authService.isLoggedIn();
  }
  isAdmin(){
    if(this.role==="admin"){
      return true;
    }
    else{
      return false;
    }
  }

  isCustomer(){
    return this.role==="customer";
  }

  isCourier(){
    return this.role==="courier";
  }

  logout(){
    this.authService.logout();
    this.userId=null;
    this.role=null;
  }






}
