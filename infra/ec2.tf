resource "aws_instance" "wap" {
  ami = "${var.ami_id}"
  instance_type = "${var.instance_type}"
  key_name = "${var.ec2_key_name}"
  associate_public_ip_address = true
  root_block_device = {
    volume_type = "${var.volume_type}"
    volume_size = "${var.volume_size}"
    delete_on_termination = true
  }
  tags {
    Name = "${var.instance_name}"
    env = "production"
    role = "wap"
  }
}
