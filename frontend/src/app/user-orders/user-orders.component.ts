import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { OrderService } from '../order.service';
import { AuthService } from '../auth.service';
import { jwtDecode } from "jwt-decode";
@Component({
  selector: 'app-user-orders',
  templateUrl: './user-orders.component.html',
  styleUrls: ['./user-orders.component.css']
})
export class UserOrdersComponent implements OnInit {
  userId: number= 0;
  orders: any[] = [];
  errorMessage: string = '';
  role:string='';
  constructor(
    /*private route: ActivatedRoute,*/ private orderService: OrderService,private authService: AuthService) {
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
    isCus(){
      if(this.role==="customer"){
        return true;
      }
      else{
        return false;
      }
  
    }

  ngOnInit(): void {
    /* Retrieve user ID from the URL parameter
    this.userId = +this.route.snapshot.paramMap.get('id')!;
    */
    
    
    // Get orders for the current user
    this.orderService.getUserOrders(this.userId).subscribe({
      next: (data) => {
        this.orders = data; // Successfully received orders
      },
      error: (error) => {
        this.errorMessage = 'There was an error retrieving orders.'; // Handle error
        console.error(error); // Log the error for debugging
      },
      complete: () => {
        console.log('User orders fetch complete'); // Optionally log completion
      }
    });
  }
}
//activated route w el snapshot l fo2 knt hst5dmm lw knt b5ly l id yt7t fil path men abl l 5atwa de fa hwa ye retrieve it mel url
//lw kan mktoob msln /orders/users/2 fa kan haya5od l 2
