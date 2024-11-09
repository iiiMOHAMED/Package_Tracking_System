import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AuthService } from 'src/app/auth.service';
import { jwtDecode } from "jwt-decode";

@Component({
  selector: 'app-orderCreation',
  templateUrl: './orderCreation.component.html',
  styleUrls: ['./orderCreation.component.css']
})
export class orderCreationComponent {
  pickupLocation: string = '';
  dropOffLocation: string = '';
  packageDetails: string = '';
  deliveryTime: string = '';
  private apiUrl = 'http://localhost:8080/orders';
  userId: number | null = null;
  role:string=" ";

  constructor(private http: HttpClient, private authService: AuthService) {
    const token = this.authService.getToken();
    if (token) {
      const decodedToken: any = jwtDecode(token);
      this.userId = decodedToken.sub; 
      this.role=decodedToken.role;
      console.log("test1:"+this.role)
      console.log("test1:"+this.role)
    }
    else{ 
      console.log("test2:"+this.role)
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

  onSubmit() {
    const order = {
      pickupLocation: this.pickupLocation,
      dropOffLocation: this.dropOffLocation,
      packageDetails: this.packageDetails,
      deliveryTime: this.deliveryTime,
      user_id: this.userId, // Include the user ID
    };

    this.http.post(this.apiUrl, order).subscribe({
      next: () => alert('Order created successfully!'),
      error: (err) => alert('Failed to create order: ' + err.error),
    });
  }
}


/*<!--type Order struct {
	OrderNumber     string `json:"orderNumber"`
	PickupLocation  string `json:"pickupLocation"`
	DropOffLocation string `json:"dropOffLocation"`
	PackageDetails  string `json:"packageDetails"`
	DeliveryTime    string `json:"deliveryTime"`
}
-->*/