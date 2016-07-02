variable "ami_id" {
  default = "ami-374db956"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "ec2_key_name" {
  default = "party"
}

variable "instance_name" {
  default = "party-wap01"
}

variable "volume_type" {
  default = "gp2"
}

variable "volume_size" {
  default = 8
}
