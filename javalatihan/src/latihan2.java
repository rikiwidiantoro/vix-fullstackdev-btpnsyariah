//public class latihan2 {
//}
class Laptopp {
    String hidupkanLaptop(String pemilik, String merk) {
        return "Hidupkan Laptop "+merk+" milik "+pemilik;
    }

    String matikanLaptop(String pemilik, String merk) {
        return "Matikan Laptop "+merk+" milik "+pemilik;
    }

    void restartLaptopp() {
        System.out.println(matikanLaptop("Riki Widiantoro", "MSI Modern 15"));
        System.out.println(hidupkanLaptop("Riki Widiantoro", "MSI Modern 15"));
    }
}

class BelajarJava2 {
    public static void main(String args[]) {
        Laptopp laptopRiki = new Laptopp();
        laptopRiki.restartLaptopp();
    }
}