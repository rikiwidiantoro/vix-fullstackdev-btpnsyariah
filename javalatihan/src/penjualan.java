public class penjualan {
    String namaPedagang(String pedagang) {
        return "Nama Pedagang : "+ pedagang;
    }
    String harga(int harga) {
        return "Harga Barang : "+ harga;
    }
    String beli(int beli) {
        return "Jumlah yang dibeli : "+ beli;
    }
    String garis(String garis) {
        return garis;
    }
    String total(String total) {
        return "Total Harga : "+ total;
    }
    int harga = 950;
    int beli = 115;
    int total = harga*beli;
    String cekJumlah() {
        return "Total Harga : "+total;
    }

    void sate() {
        System.out.println(namaPedagang("Pak Kumis"));
        System.out.println(harga(harga));
        System.out.println(beli(beli));
        System.out.println(garis("_________________________"));
        System.out.println(cekJumlah());
    }
}

class Tampilkan {
    public static void main(String args[]) {
        penjualan jual = new penjualan();
        jual.sate();
    }
}
